package main

// Openjph - Open-source implementation of JPEG2000 Part-15 (or JPH or HTJ2K)
// Homepage: https://github.com/aous72/OpenJPH

import (
	"fmt"
	
	"os/exec"
)

func installOpenjph() {
	// Método 1: Descargar y extraer .tar.gz
	openjph_tar_url := "https://github.com/aous72/OpenJPH/archive/refs/tags/0.15.0.tar.gz"
	openjph_cmd_tar := exec.Command("curl", "-L", openjph_tar_url, "-o", "package.tar.gz")
	err := openjph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openjph_zip_url := "https://github.com/aous72/OpenJPH/archive/refs/tags/0.15.0.zip"
	openjph_cmd_zip := exec.Command("curl", "-L", openjph_zip_url, "-o", "package.zip")
	err = openjph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openjph_bin_url := "https://github.com/aous72/OpenJPH/archive/refs/tags/0.15.0.bin"
	openjph_cmd_bin := exec.Command("curl", "-L", openjph_bin_url, "-o", "binary.bin")
	err = openjph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openjph_src_url := "https://github.com/aous72/OpenJPH/archive/refs/tags/0.15.0.src.tar.gz"
	openjph_cmd_src := exec.Command("curl", "-L", openjph_src_url, "-o", "source.tar.gz")
	err = openjph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openjph_cmd_direct := exec.Command("./binary")
	err = openjph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
}
