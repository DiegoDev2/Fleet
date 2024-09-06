package main

// Uhdm - Universal Hardware Data Model, modeling of the SystemVerilog Object Model
// Homepage: https://github.com/chipsalliance/UHDM

import (
	"fmt"
	
	"os/exec"
)

func installUhdm() {
	// Método 1: Descargar y extraer .tar.gz
	uhdm_tar_url := "https://github.com/chipsalliance/UHDM/archive/refs/tags/v1.84.tar.gz"
	uhdm_cmd_tar := exec.Command("curl", "-L", uhdm_tar_url, "-o", "package.tar.gz")
	err := uhdm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uhdm_zip_url := "https://github.com/chipsalliance/UHDM/archive/refs/tags/v1.84.zip"
	uhdm_cmd_zip := exec.Command("curl", "-L", uhdm_zip_url, "-o", "package.zip")
	err = uhdm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uhdm_bin_url := "https://github.com/chipsalliance/UHDM/archive/refs/tags/v1.84.bin"
	uhdm_cmd_bin := exec.Command("curl", "-L", uhdm_bin_url, "-o", "binary.bin")
	err = uhdm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uhdm_src_url := "https://github.com/chipsalliance/UHDM/archive/refs/tags/v1.84.src.tar.gz"
	uhdm_cmd_src := exec.Command("curl", "-L", uhdm_src_url, "-o", "source.tar.gz")
	err = uhdm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uhdm_cmd_direct := exec.Command("./binary")
	err = uhdm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: capnp")
	exec.Command("latte", "install", "capnp").Run()
}
