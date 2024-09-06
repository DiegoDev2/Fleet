package main

// Pngquant - PNG image optimizing utility
// Homepage: https://pngquant.org/

import (
	"fmt"
	
	"os/exec"
)

func installPngquant() {
	// Método 1: Descargar y extraer .tar.gz
	pngquant_tar_url := "https://static.crates.io/crates/pngquant/pngquant-3.0.3.crate"
	pngquant_cmd_tar := exec.Command("curl", "-L", pngquant_tar_url, "-o", "package.tar.gz")
	err := pngquant_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pngquant_zip_url := "https://static.crates.io/crates/pngquant/pngquant-3.0.3.crate"
	pngquant_cmd_zip := exec.Command("curl", "-L", pngquant_zip_url, "-o", "package.zip")
	err = pngquant_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pngquant_bin_url := "https://static.crates.io/crates/pngquant/pngquant-3.0.3.crate"
	pngquant_cmd_bin := exec.Command("curl", "-L", pngquant_bin_url, "-o", "binary.bin")
	err = pngquant_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pngquant_src_url := "https://static.crates.io/crates/pngquant/pngquant-3.0.3.crate"
	pngquant_cmd_src := exec.Command("curl", "-L", pngquant_src_url, "-o", "source.tar.gz")
	err = pngquant_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pngquant_cmd_direct := exec.Command("./binary")
	err = pngquant_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
}
