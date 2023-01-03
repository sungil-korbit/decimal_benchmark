package src;

import org.openjdk.jmh.annotations.*;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.List;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

@BenchmarkMode(Mode.SingleShotTime)
@OutputTimeUnit(TimeUnit.NANOSECONDS)
@Fork(value = 1, jvmArgs = {"-Xms4g", "-Xmx4g"})
public class Bench {
    public void decimalOp(DecimalOp op, BigDecimal x, BigDecimal y) {
        switch (op) {
            case ADD: x.add(y); break;
            case SUB: x.subtract(y); break;
            case MUL: x.multiply(y); break;
            case DIV: x.divide(y, 2, RoundingMode.CEILING);break;
            default:
        }
    }

    public void foldOp(DecimalOp op, BigDecimal init, int count) {
        List<BigDecimal> values = IntStream.rangeClosed(1, count)
                .mapToObj(i -> new BigDecimal(i))
                .collect(Collectors.toList());

        BigDecimal acc = init;

        switch (op) {
            case ADD:
                for (BigDecimal v : values) acc = acc.add(v);
                break;
            case SUB:
                for (BigDecimal v : values) acc = acc.subtract(v);
                break;
            case MUL:
                for (BigDecimal v : values) acc = acc.multiply(v);
                break;
            case DIV:
                for (BigDecimal v : values) acc = acc.divide(v, 2 , RoundingMode.CEILING);break;
            default:
        }
    }

//    @TearDown
//    public void tearDown() {
//        System.gc();
//    }
    /* ADD */
    @Benchmark
    public void addSelf() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("2.01"));}
    @Benchmark
    public void addSimple() {decimalOp(DecimalOp.ADD, new BigDecimal("2"), new BigDecimal("1"));}
    @Benchmark
    public void addOne() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("1"));}
    @Benchmark
    public void addTwo() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("2"));}
    @Benchmark
    public void addOneHundred() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("100"));}
    @Benchmark
    public void addPointZeroOne() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("0.01"));}
    @Benchmark
    public void addNegativePointFive() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("-0.5"));}
    @Benchmark
    public void addPi() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("3.1415926535897932384626433832"));}
    @Benchmark
    public void addNegativePi() {decimalOp(DecimalOp.ADD, new BigDecimal("2.01"), new BigDecimal("-3.1415926535897932384626433832"));}
    @Benchmark
    public void add10k() {foldOp(DecimalOp.ADD, new BigDecimal("0"), 10000);}

    /* SUB */
    @Benchmark
    public void subSelf() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("2.01"));}
    @Benchmark
    public void subSimple() {decimalOp(DecimalOp.SUB, new BigDecimal("2"), new BigDecimal("1"));}
    @Benchmark
    public void subOne() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("1"));}
    @Benchmark
    public void subTwo() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("2"));}
    @Benchmark
    public void subOneHundred() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("100"));}
    @Benchmark
    public void subPointZeroOne() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("0.01"));}
    @Benchmark
    public void subNegativePointFive() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("-0.5"));}
    @Benchmark
    public void subPi() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("3.1415926535897932384626433832"));}
    @Benchmark
    public void subNegativePi() {decimalOp(DecimalOp.SUB, new BigDecimal("2.01"), new BigDecimal("-3.1415926535897932384626433832"));}
    @Benchmark
    public void sub10k() {foldOp(DecimalOp.SUB, new BigDecimal("5000000"), 10000);}


    /* MUL */
    @Benchmark
    public void mulSelf() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("2.01"));}
    @Benchmark
    public void mulSimple() {decimalOp(DecimalOp.MUL, new BigDecimal("2"), new BigDecimal("1"));}
    @Benchmark
    public void mulOne() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("1"));}
    @Benchmark
    public void mulTwo() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("2"));}
    @Benchmark
    public void mulOneHundred() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("100"));}
    @Benchmark
    public void mulPointZeroOne() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("0.01"));}
    @Benchmark
    public void mulNegativePointFive() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("-0.5"));}
    @Benchmark
    public void mulPi() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("3.1415926535897932384626433832"));}
    @Benchmark
    public void mulNegativePi() {decimalOp(DecimalOp.MUL, new BigDecimal("2.01"), new BigDecimal("-3.1415926535897932384626433832"));}
    @Benchmark
    public void mul10k() {foldOp(DecimalOp.MUL, new BigDecimal("1.1"), 25);}

    /* DIV */
    @Benchmark
    public void divSelf() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("2.01"));}
    @Benchmark
    public void divSimple() {decimalOp(DecimalOp.DIV, new BigDecimal("2"), new BigDecimal("1"));}
    @Benchmark
    public void divOne() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("1"));}
    @Benchmark
    public void divTwo() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("2"));}
    @Benchmark
    public void divOneHundred() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("100"));}
    @Benchmark
    public void divPointZeroOne() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("0.01"));}
    @Benchmark
    public void divNegativePointFive() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("-0.5"));}
    @Benchmark
    public void divPi() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("3.1415926535897932384626433832"));}
    @Benchmark
    public void divNegativePi() {decimalOp(DecimalOp.DIV, new BigDecimal("2.01"), new BigDecimal("-3.1415926535897932384626433832"));}
    @Benchmark
    public void div10k() {foldOp(DecimalOp.DIV, new BigDecimal("1000000000"), 10000);}
}
