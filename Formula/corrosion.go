package main

// Corrosion - Easy Rust and C/C++ Integration
// Homepage: https://github.com/corrosion-rs/corrosion

import (
	"fmt"
	
	"os/exec"
)

func installCorrosion() {
	// Método 1: Descargar y extraer .tar.gz
	corrosion_tar_url := "https://github.com/corrosion-rs/corrosion/archive/refs/tags/v0.5.tar.gz"
	corrosion_cmd_tar := exec.Command("curl", "-L", corrosion_tar_url, "-o", "package.tar.gz")
	err := corrosion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	corrosion_zip_url := "https://github.com/corrosion-rs/corrosion/archive/refs/tags/v0.5.zip"
	corrosion_cmd_zip := exec.Command("curl", "-L", corrosion_zip_url, "-o", "package.zip")
	err = corrosion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	corrosion_bin_url := "https://github.com/corrosion-rs/corrosion/archive/refs/tags/v0.5.bin"
	corrosion_cmd_bin := exec.Command("curl", "-L", corrosion_bin_url, "-o", "binary.bin")
	err = corrosion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	corrosion_src_url := "https://github.com/corrosion-rs/corrosion/archive/refs/tags/v0.5.src.tar.gz"
	corrosion_cmd_src := exec.Command("curl", "-L", corrosion_src_url, "-o", "source.tar.gz")
	err = corrosion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	corrosion_cmd_direct := exec.Command("./binary")
	err = corrosion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
