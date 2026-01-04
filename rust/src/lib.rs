use std::{mem::size_of, os::raw::c_int};

/// Native Rust implementation
pub mod core {
    pub fn add(a: i32, b: i32) -> i32 {
        a + b
    }
}

/// FFI adapter layer - converts between C and Rust types
#[unsafe(no_mangle)]
pub extern "C" fn add(a: c_int, b: c_int) -> c_int {
    // Compile-time assertion that c_int is i32 on this platform
    const _: () = assert!(size_of::<c_int>() == size_of::<i32>());

    core::add(a as i32, b as i32) as c_int
}
