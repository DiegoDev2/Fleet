package main

// Netris - Networked variant of tetris
// Homepage: https://packages.debian.org/sid/netris

import (
	"fmt"
	
	"os/exec"
)

func installNetris() {
	// Método 1: Descargar y extraer .tar.gz
	netris_tar_url := "https://deb.debian.org/debian/pool/main/n/netris/netris_0.52.orig.tar.gz"
	netris_cmd_tar := exec.Command("curl", "-L", netris_tar_url, "-o", "package.tar.gz")
	err := netris_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netris_zip_url := "https://deb.debian.org/debian/pool/main/n/netris/netris_0.52.orig.zip"
	netris_cmd_zip := exec.Command("curl", "-L", netris_zip_url, "-o", "package.zip")
	err = netris_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netris_bin_url := "https://deb.debian.org/debian/pool/main/n/netris/netris_0.52.orig.bin"
	netris_cmd_bin := exec.Command("curl", "-L", netris_bin_url, "-o", "binary.bin")
	err = netris_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netris_src_url := "https://deb.debian.org/debian/pool/main/n/netris/netris_0.52.orig.src.tar.gz"
	netris_cmd_src := exec.Command("curl", "-L", netris_src_url, "-o", "source.tar.gz")
	err = netris_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netris_cmd_direct := exec.Command("./binary")
	err = netris_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
