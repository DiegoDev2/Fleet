package main

// Nacl - Network communication, encryption, decryption, signatures library
// Homepage: https://nacl.cr.yp.to/

import (
	"fmt"
	
	"os/exec"
)

func installNacl() {
	// Método 1: Descargar y extraer .tar.gz
	nacl_tar_url := "https://hyperelliptic.org/nacl/nacl-20110221.tar.bz2"
	nacl_cmd_tar := exec.Command("curl", "-L", nacl_tar_url, "-o", "package.tar.gz")
	err := nacl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nacl_zip_url := "https://hyperelliptic.org/nacl/nacl-20110221.tar.bz2"
	nacl_cmd_zip := exec.Command("curl", "-L", nacl_zip_url, "-o", "package.zip")
	err = nacl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nacl_bin_url := "https://hyperelliptic.org/nacl/nacl-20110221.tar.bz2"
	nacl_cmd_bin := exec.Command("curl", "-L", nacl_bin_url, "-o", "binary.bin")
	err = nacl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nacl_src_url := "https://hyperelliptic.org/nacl/nacl-20110221.tar.bz2"
	nacl_cmd_src := exec.Command("curl", "-L", nacl_src_url, "-o", "source.tar.gz")
	err = nacl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nacl_cmd_direct := exec.Command("./binary")
	err = nacl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
