package main

// Veilid - Peer-to-peer network for easily sharing various kinds of data
// Homepage: https://veilid.com/

import (
	"fmt"
	
	"os/exec"
)

func installVeilid() {
	// Método 1: Descargar y extraer .tar.gz
	veilid_tar_url := "https://gitlab.com/veilid/veilid/-/archive/v0.3.4/veilid-v0.3.4.tar.bz2"
	veilid_cmd_tar := exec.Command("curl", "-L", veilid_tar_url, "-o", "package.tar.gz")
	err := veilid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	veilid_zip_url := "https://gitlab.com/veilid/veilid/-/archive/v0.3.4/veilid-v0.3.4.tar.bz2"
	veilid_cmd_zip := exec.Command("curl", "-L", veilid_zip_url, "-o", "package.zip")
	err = veilid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	veilid_bin_url := "https://gitlab.com/veilid/veilid/-/archive/v0.3.4/veilid-v0.3.4.tar.bz2"
	veilid_cmd_bin := exec.Command("curl", "-L", veilid_bin_url, "-o", "binary.bin")
	err = veilid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	veilid_src_url := "https://gitlab.com/veilid/veilid/-/archive/v0.3.4/veilid-v0.3.4.tar.bz2"
	veilid_cmd_src := exec.Command("curl", "-L", veilid_src_url, "-o", "source.tar.gz")
	err = veilid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	veilid_cmd_direct := exec.Command("./binary")
	err = veilid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
