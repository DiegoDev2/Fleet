package main

// Reop - Encrypted keypair management
// Homepage: https://flak.tedunangst.com/post/reop

import (
	"fmt"
	
	"os/exec"
)

func installReop() {
	// Método 1: Descargar y extraer .tar.gz
	reop_tar_url := "https://flak.tedunangst.com/files/reop-2.1.1.tgz"
	reop_cmd_tar := exec.Command("curl", "-L", reop_tar_url, "-o", "package.tar.gz")
	err := reop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reop_zip_url := "https://flak.tedunangst.com/files/reop-2.1.1.tgz"
	reop_cmd_zip := exec.Command("curl", "-L", reop_zip_url, "-o", "package.zip")
	err = reop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reop_bin_url := "https://flak.tedunangst.com/files/reop-2.1.1.tgz"
	reop_cmd_bin := exec.Command("curl", "-L", reop_bin_url, "-o", "binary.bin")
	err = reop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reop_src_url := "https://flak.tedunangst.com/files/reop-2.1.1.tgz"
	reop_cmd_src := exec.Command("curl", "-L", reop_src_url, "-o", "source.tar.gz")
	err = reop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reop_cmd_direct := exec.Command("./binary")
	err = reop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
}
