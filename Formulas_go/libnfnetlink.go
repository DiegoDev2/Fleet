package main

// Libnfnetlink - Low-level library for netfilter related communication
// Homepage: https://www.netfilter.org/projects/libnfnetlink

import (
	"fmt"
	
	"os/exec"
)

func installLibnfnetlink() {
	// Método 1: Descargar y extraer .tar.gz
	libnfnetlink_tar_url := "https://www.netfilter.org/projects/libnfnetlink/files/libnfnetlink-1.0.2.tar.bz2"
	libnfnetlink_cmd_tar := exec.Command("curl", "-L", libnfnetlink_tar_url, "-o", "package.tar.gz")
	err := libnfnetlink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnfnetlink_zip_url := "https://www.netfilter.org/projects/libnfnetlink/files/libnfnetlink-1.0.2.tar.bz2"
	libnfnetlink_cmd_zip := exec.Command("curl", "-L", libnfnetlink_zip_url, "-o", "package.zip")
	err = libnfnetlink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnfnetlink_bin_url := "https://www.netfilter.org/projects/libnfnetlink/files/libnfnetlink-1.0.2.tar.bz2"
	libnfnetlink_cmd_bin := exec.Command("curl", "-L", libnfnetlink_bin_url, "-o", "binary.bin")
	err = libnfnetlink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnfnetlink_src_url := "https://www.netfilter.org/projects/libnfnetlink/files/libnfnetlink-1.0.2.tar.bz2"
	libnfnetlink_cmd_src := exec.Command("curl", "-L", libnfnetlink_src_url, "-o", "source.tar.gz")
	err = libnfnetlink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnfnetlink_cmd_direct := exec.Command("./binary")
	err = libnfnetlink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
