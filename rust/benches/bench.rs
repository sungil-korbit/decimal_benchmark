#![feature(test)]

extern crate test;

use core::str::FromStr;
use rust_decimal::Decimal;
use test::Bencher;

macro_rules! bench_decimal_op {
    ($name:ident, $op:tt, $x:expr, $y:expr) => {
        #[bench]
        fn $name(b: &mut Bencher) {
            let x = Decimal::from_str($x).unwrap();
            let y = Decimal::from_str($y).unwrap();

            b.iter(|| {
                let result = x $op y;
                ::test::black_box(result);
            });
        }
    }
}


macro_rules! bench_fold_op {
    ($name:ident, $op:tt, $init:expr, $count:expr) => {
        #[bench]
        fn $name(b: &mut ::test::Bencher) {
            fn fold(values: &[Decimal]) -> Decimal {
                let mut acc: Decimal = $init.into();
                for value in values {
                    acc = acc $op value;
                }
                acc
            }

            let values: Vec<Decimal> = test::black_box((1..$count).map(|i| i.into()).collect());
            b.iter(|| {
                let result = fold(&values);
                ::test::black_box(result);
            });
        }
    }
}

/* Add */
bench_decimal_op!(add_self, +, "2.01", "2.01");
bench_decimal_op!(add_simple, +, "2", "1");
bench_decimal_op!(add_one, +, "2.01", "1");
bench_decimal_op!(add_two, +, "2.01", "2");
bench_decimal_op!(add_one_hundred, +, "2.01", "100");
bench_decimal_op!(add_point_zero_one, +, "2.01", "0.01");
bench_decimal_op!(add_negative_point_five, +, "2.01", "-0.5");
bench_decimal_op!(add_pi, +, "2.01", "3.1415926535897932384626433832");
bench_decimal_op!(add_negative_pi, +, "2.01", "-3.1415926535897932384626433832");

bench_fold_op!(add_10k, +, 0, 10_000);

/* Sub */
bench_decimal_op!(sub_self, -, "2.01", "2.01");
bench_decimal_op!(sub_simple, -, "2", "1");
bench_decimal_op!(sub_one, -, "2.01", "1");
bench_decimal_op!(sub_two, -, "2.01", "2");
bench_decimal_op!(sub_one_hundred, -, "2.01", "100");
bench_decimal_op!(sub_point_zero_one, -, "2.01", "0.01");
bench_decimal_op!(sub_negative_point_five, -, "2.01", "-0.5");
bench_decimal_op!(sub_pi, -, "2.01", "3.1415926535897932384626433832");
bench_decimal_op!(sub_negative_pi, -, "2.01", "-3.1415926535897932384626433832");

bench_fold_op!(sub_10k, -, 5_000_000, 10_000);

/* Mul */
bench_decimal_op!(mul_one, *, "2.01", "1");
bench_decimal_op!(mul_two, *, "2.01", "2");
bench_decimal_op!(mul_one_hundred, *, "2.01", "100");
bench_decimal_op!(mul_point_zero_one, *, "2.01", "0.01");
bench_decimal_op!(mul_negative_point_five, *, "2.01", "-0.5");
bench_decimal_op!(mul_pi, *, "2.01", "3.1415926535897932384626433832");
bench_decimal_op!(mul_negative_pi, *, "2.01", "-3.1415926535897932384626433832");

bench_fold_op!(mul_25, *, Decimal::from_str("1.1").unwrap(), 25);

/* Div */
bench_decimal_op!(div_one, /, "2.01", "1");
bench_decimal_op!(div_two, /, "2.01", "2");
bench_decimal_op!(div_one_hundred, /, "2.01", "100");
bench_decimal_op!(div_point_zero_one, /, "2.01", "0.01");
bench_decimal_op!(div_negative_point_five, /, "2.01", "-0.5");
bench_decimal_op!(div_pi, /, "2.01", "3.1415926535897932384626433832");
bench_decimal_op!(div_negative_pi, /, "2.01", "-3.1415926535897932384626433832");
bench_decimal_op!(div_no_underflow, /, "1.02343545345", "0.35454343453");
bench_fold_op!(div_10k, /, Decimal::MAX, 10_000);
bench_fold_op!(rem_10k, %, Decimal::MAX, 10_000);