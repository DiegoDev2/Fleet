package main

// Aamath - Renders mathematical expressions as ASCII art
// Homepage: http://fuse.superglue.se/aamath/

import (
	"fmt"
	
	"os/exec"
)

func installAamath() {
	// Método 1: Descargar y extraer .tar.gz
	aamath_tar_url := "http://fuse.superglue.se/aamath/aamath-0.3.tar.gz"
	aamath_cmd_tar := exec.Command("curl", "-L", aamath_tar_url, "-o", "package.tar.gz")
	err := aamath_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aamath_zip_url := "http://fuse.superglue.se/aamath/aamath-0.3.zip"
	aamath_cmd_zip := exec.Command("curl", "-L", aamath_zip_url, "-o", "package.zip")
	err = aamath_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aamath_bin_url := "http://fuse.superglue.se/aamath/aamath-0.3.bin"
	aamath_cmd_bin := exec.Command("curl", "-L", aamath_bin_url, "-o", "binary.bin")
	err = aamath_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aamath_src_url := "http://fuse.superglue.se/aamath/aamath-0.3.src.tar.gz"
	aamath_cmd_src := exec.Command("curl", "-L", aamath_src_url, "-o", "source.tar.gz")
	err = aamath_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aamath_cmd_direct := exec.Command("./binary")
	err = aamath_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
