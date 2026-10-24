# Concurrency Patterns in Go

## 1. Pipeline

[`patterns/pipeline`](patterns/pipeline/pattern.go)

Последовательная обработка данных через каналы.

```text
Generate -> Square -> FilterEven -> Result
```

## 2. Fan-out / Fan-in

[`patterns/fan_in-fan_out`](patterns/fan_in-fan_out/pattern.go)

Распределение задач между горутинами и сбор результатов в один канал.

```text
         +-> Worker 1 -+
Input ---+-> Worker 2 -+---> Result
         +-> Worker 3 -+
```

## 3. Producer–Consumer

[`patterns/producer-consumer`](patterns/producer-consumer/pattern.go)

Производители отправляют данные в общий канал, потребители обрабатывают их.

```text
Producers -> Channel -> Consumers -> Result
```

## 4. Worker Pull

[`patterns/worker_pull`](patterns/worker_pull/pattern.go)

Воркеры самостоятельно забирают задачи из общего канала.

```text
                +-> Worker 1 -+
Tasks -> Queue -+-> Worker 2 -+-> Result
                +-> Worker 3 -+
```

## 5. Semaphore

[`patterns/semophore`](patterns/semophore/pattern.go)

Ограничение количества одновременно выполняемых задач через буферизированный канал.

```text
Tasks -> Semaphore (limit) -> Execution
```

## 6. Singleflight

[`patterns/singleflight`](patterns/singleflight/pattern.go)

Одновременные вызовы с одним ключом используют результат одного выполнения.

```text
Do("key") -+
Do("key") -+-> fn() -> Shared result
Do("key") -+
```

## 7. Cancellation & Backpressure

[`patterns/cancellation-backpressure`](patterns/cancellation-backpressure/pattern.go)

`context.Context` отменяет обработку; медленный получатель блокирует отправку.

```text
Producer -> Channel -> Consumer
    |          |
    +-- ctx.Done() -> Stop
```
