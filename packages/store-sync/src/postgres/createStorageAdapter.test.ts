import { beforeEach, describe, expect, it } from "vitest";
import { DefaultLogger, eq } from "drizzle-orm";
import { drizzle } from "drizzle-orm/postgres-js";
import postgres from "postgres";
import { Hex, RpcLog, createPublicClient, decodeEventLog, formatLog, http } from "viem";
import { foundry } from "viem/chains";
import { PostgresStorageAdapter, createStorageAdapter } from "./createStorageAdapter";
import { groupLogsByBlockNumber } from "@latticexyz/block-logs-stream";
import { storeEventsAbi } from "@latticexyz/store";
import { StoreEventsLog } from "../common";
import worldRpcLogs from "../../../../test-data/world-logs.json";
import { resourceToHex } from "@latticexyz/common";

const blocks = groupLogsByBlockNumber(
  worldRpcLogs.map((log) => {
    const { eventName, args } = decodeEventLog({
      abi: storeEventsAbi,
      data: log.data as Hex,
      topics: log.topics as [Hex, ...Hex[]],
      strict: true,
    });
    return formatLog(log as any as RpcLog, { args, eventName: eventName as string }) as StoreEventsLog;
  })
);

describe("createStorageAdapter", async () => {
  const db = drizzle(postgres(process.env.DATABASE_URL!), {
    logger: new DefaultLogger(),
  });

  const publicClient = createPublicClient({
    chain: foundry,
    transport: http(),
  });

  let storageAdapter: PostgresStorageAdapter;

  beforeEach(async () => {
    storageAdapter = await createStorageAdapter({ database: db, publicClient });
    return storageAdapter.cleanUp;
  });

  it("should create tables and data from block log", async () => {
    for (const block of blocks) {
      await storageAdapter.storageAdapter(block);
    }

    expect(await db.select().from(storageAdapter.tables.configTable)).toMatchInlineSnapshot(`
      [
        {
          "blockNumber": 20n,
          "chainId": 31337,
          "version": "0.0.4",
        },
      ]
    `);

    expect(
      await db
        .select()
        .from(storageAdapter.tables.recordsTable)
        .where(
          eq(
            storageAdapter.tables.recordsTable.tableId,
            resourceToHex({ type: "table", namespace: "", name: "NumberList" })
          )
        )
    ).toMatchInlineSnapshot(`
      [
        {
          "address": "0x2964aF56c8aACdE425978a28b018956D21cF50f0",
          "blockNumber": 20n,
          "dynamicData": "0x000001a400000045",
          "encodedLengths": "0x0000000000000000000000000000000000000000000000000800000000000008",
          "isDeleted": false,
          "key0": null,
          "key1": null,
          "keyBytes": "0x",
          "logIndex": 1,
          "staticData": null,
          "tableId": "0x746200000000000000000000000000004e756d6265724c697374000000000000",
        },
      ]
    `);

    await storageAdapter.cleanUp();
  }, 15_000);
});
