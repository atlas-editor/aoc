mod days;
mod utils;

fn main() {
    // benchmark_all!(
    //     d01, d02, d03, d04, d05, d06, d07, d08, d09, d10, d11, d12, d13, d14, d15, d16, d17, d18
    // );
    run!(d11);
    benchmark_all!(d11);
}
