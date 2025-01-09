use std::fs;
use std::path::Path;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let out_dir = "src/proto_generated";
    if !Path::new(out_dir).exists() {
        fs::create_dir_all(out_dir)?;
    }

    tonic_build::configure()
        .out_dir(out_dir) // Specify the output directory
        .compile_protos(&["proto/facade.proto"], &["proto"])?; // Adjust paths to your proto files
    println!("cargo:rerun-if-changed=proto/facade.proto");
    Ok(())
}
