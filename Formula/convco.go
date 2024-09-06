package main

// Convco - Conventional commits, changelog, versioning, validation
// Homepage: https://convco.github.io

import (
	"fmt"
	
	"os/exec"
)

func installConvco() {
	// Método 1: Descargar y extraer .tar.gz
	convco_tar_url := "https://github.com/convco/convco/archive/refs/tags/v0.5.2.tar.gz"
	convco_cmd_tar := exec.Command("curl", "-L", convco_tar_url, "-o", "package.tar.gz")
	err := convco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	convco_zip_url := "https://github.com/convco/convco/archive/refs/tags/v0.5.2.zip"
	convco_cmd_zip := exec.Command("curl", "-L", convco_zip_url, "-o", "package.zip")
	err = convco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	convco_bin_url := "https://github.com/convco/convco/archive/refs/tags/v0.5.2.bin"
	convco_cmd_bin := exec.Command("curl", "-L", convco_bin_url, "-o", "binary.bin")
	err = convco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	convco_src_url := "https://github.com/convco/convco/archive/refs/tags/v0.5.2.src.tar.gz"
	convco_cmd_src := exec.Command("curl", "-L", convco_src_url, "-o", "source.tar.gz")
	err = convco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	convco_cmd_direct := exec.Command("./binary")
	err = convco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
}
