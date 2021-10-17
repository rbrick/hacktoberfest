
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
        _ => {
            let (a, b) = (0, 1);
            for _ in 0..i {
                let c = a + b;
                a = b;
                b = c;
                // sadly, this isn't stable
                // (a, b) = (b, a + b)
            }
            a


        }
    }
}
