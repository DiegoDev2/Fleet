package main

// Crabz - Like pigz, but in Rust
// Homepage: https://github.com/sstadick/crabz

import (
	"fmt"
	
	"os/exec"
)

func installCrabz() {
	// Método 1: Descargar y extraer .tar.gz
	crabz_tar_url := "https://github.com/sstadick/crabz/archive/refs/tags/v0.10.0.tar.gz"
	crabz_cmd_tar := exec.Command("curl", "-L", crabz_tar_url, "-o", "package.tar.gz")
	err := crabz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crabz_zip_url := "https://github.com/sstadick/crabz/archive/refs/tags/v0.10.0.zip"
	crabz_cmd_zip := exec.Command("curl", "-L", crabz_zip_url, "-o", "package.zip")
	err = crabz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crabz_bin_url := "https://github.com/sstadick/crabz/archive/refs/tags/v0.10.0.bin"
	crabz_cmd_bin := exec.Command("curl", "-L", crabz_bin_url, "-o", "binary.bin")
	err = crabz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crabz_src_url := "https://github.com/sstadick/crabz/archive/refs/tags/v0.10.0.src.tar.gz"
	crabz_cmd_src := exec.Command("curl", "-L", crabz_src_url, "-o", "source.tar.gz")
	err = crabz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crabz_cmd_direct := exec.Command("./binary")
	err = crabz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
