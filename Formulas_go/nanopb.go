package main

// Nanopb - C library for encoding and decoding Protocol Buffer messages
// Homepage: https://jpa.kapsi.fi/nanopb/docs/index.html

import (
	"fmt"
	
	"os/exec"
)

func installNanopb() {
	// Método 1: Descargar y extraer .tar.gz
	nanopb_tar_url := "https://jpa.kapsi.fi/nanopb/download/nanopb-0.4.8.tar.gz"
	nanopb_cmd_tar := exec.Command("curl", "-L", nanopb_tar_url, "-o", "package.tar.gz")
	err := nanopb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nanopb_zip_url := "https://jpa.kapsi.fi/nanopb/download/nanopb-0.4.8.zip"
	nanopb_cmd_zip := exec.Command("curl", "-L", nanopb_zip_url, "-o", "package.zip")
	err = nanopb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nanopb_bin_url := "https://jpa.kapsi.fi/nanopb/download/nanopb-0.4.8.bin"
	nanopb_cmd_bin := exec.Command("curl", "-L", nanopb_bin_url, "-o", "binary.bin")
	err = nanopb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nanopb_src_url := "https://jpa.kapsi.fi/nanopb/download/nanopb-0.4.8.src.tar.gz"
	nanopb_cmd_src := exec.Command("curl", "-L", nanopb_src_url, "-o", "source.tar.gz")
	err = nanopb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nanopb_cmd_direct := exec.Command("./binary")
	err = nanopb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
