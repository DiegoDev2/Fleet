package main

// CargoZigbuild - Compile Cargo project with zig as linker
// Homepage: https://github.com/rust-cross/cargo-zigbuild

import (
	"fmt"
	
	"os/exec"
)

func installCargoZigbuild() {
	// Método 1: Descargar y extraer .tar.gz
	cargozigbuild_tar_url := "https://github.com/rust-cross/cargo-zigbuild/archive/refs/tags/v0.19.1.tar.gz"
	cargozigbuild_cmd_tar := exec.Command("curl", "-L", cargozigbuild_tar_url, "-o", "package.tar.gz")
	err := cargozigbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargozigbuild_zip_url := "https://github.com/rust-cross/cargo-zigbuild/archive/refs/tags/v0.19.1.zip"
	cargozigbuild_cmd_zip := exec.Command("curl", "-L", cargozigbuild_zip_url, "-o", "package.zip")
	err = cargozigbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargozigbuild_bin_url := "https://github.com/rust-cross/cargo-zigbuild/archive/refs/tags/v0.19.1.bin"
	cargozigbuild_cmd_bin := exec.Command("curl", "-L", cargozigbuild_bin_url, "-o", "binary.bin")
	err = cargozigbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargozigbuild_src_url := "https://github.com/rust-cross/cargo-zigbuild/archive/refs/tags/v0.19.1.src.tar.gz"
	cargozigbuild_cmd_src := exec.Command("curl", "-L", cargozigbuild_src_url, "-o", "source.tar.gz")
	err = cargozigbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargozigbuild_cmd_direct := exec.Command("./binary")
	err = cargozigbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
	fmt.Println("Instalando dependencia: zig")
	exec.Command("latte", "install", "zig").Run()
}
