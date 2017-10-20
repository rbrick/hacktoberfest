
fn main() {
    let a = fib(44);
    println!("The 44th fibonnaci number is {}",a);
    println!("Now I want a free T-shirt.");
}

fn fib(i: usize) -> usize {
    match i {
        0 => unreachable!(),
        1 => 1,
        2 => 1,
        _ => fib(i-1) + fib(i-2)
    }
}
