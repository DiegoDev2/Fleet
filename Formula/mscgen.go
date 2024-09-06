package main

// Mscgen - Parses Message Sequence Chart descriptions and produces images
// Homepage: https://www.mcternan.me.uk/mscgen/

import (
	"fmt"
	
	"os/exec"
)

func installMscgen() {
	// Método 1: Descargar y extraer .tar.gz
	mscgen_tar_url := "https://www.mcternan.me.uk/mscgen/software/mscgen-src-0.20.tar.gz"
	mscgen_cmd_tar := exec.Command("curl", "-L", mscgen_tar_url, "-o", "package.tar.gz")
	err := mscgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mscgen_zip_url := "https://www.mcternan.me.uk/mscgen/software/mscgen-src-0.20.zip"
	mscgen_cmd_zip := exec.Command("curl", "-L", mscgen_zip_url, "-o", "package.zip")
	err = mscgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mscgen_bin_url := "https://www.mcternan.me.uk/mscgen/software/mscgen-src-0.20.bin"
	mscgen_cmd_bin := exec.Command("curl", "-L", mscgen_bin_url, "-o", "binary.bin")
	err = mscgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mscgen_src_url := "https://www.mcternan.me.uk/mscgen/software/mscgen-src-0.20.src.tar.gz"
	mscgen_cmd_src := exec.Command("curl", "-L", mscgen_src_url, "-o", "source.tar.gz")
	err = mscgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mscgen_cmd_direct := exec.Command("./binary")
	err = mscgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
}
