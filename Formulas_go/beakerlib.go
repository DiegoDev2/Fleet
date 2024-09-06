package main

// Beakerlib - Shell-level integration testing library
// Homepage: https://github.com/beakerlib/beakerlib

import (
	"fmt"
	
	"os/exec"
)

func installBeakerlib() {
	// Método 1: Descargar y extraer .tar.gz
	beakerlib_tar_url := "https://github.com/beakerlib/beakerlib/archive/refs/tags/1.31.2.tar.gz"
	beakerlib_cmd_tar := exec.Command("curl", "-L", beakerlib_tar_url, "-o", "package.tar.gz")
	err := beakerlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	beakerlib_zip_url := "https://github.com/beakerlib/beakerlib/archive/refs/tags/1.31.2.zip"
	beakerlib_cmd_zip := exec.Command("curl", "-L", beakerlib_zip_url, "-o", "package.zip")
	err = beakerlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	beakerlib_bin_url := "https://github.com/beakerlib/beakerlib/archive/refs/tags/1.31.2.bin"
	beakerlib_cmd_bin := exec.Command("curl", "-L", beakerlib_bin_url, "-o", "binary.bin")
	err = beakerlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	beakerlib_src_url := "https://github.com/beakerlib/beakerlib/archive/refs/tags/1.31.2.src.tar.gz"
	beakerlib_cmd_src := exec.Command("curl", "-L", beakerlib_src_url, "-o", "source.tar.gz")
	err = beakerlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	beakerlib_cmd_direct := exec.Command("./binary")
	err = beakerlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: gnu-getopt")
exec.Command("latte", "install", "gnu-getopt")
}
