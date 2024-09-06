package main

// Mmsrip - Client for the MMS:// protocol
// Homepage: https://nbenoit.tuxfamily.org/index.php?page=MMSRIP

import (
	"fmt"
	
	"os/exec"
)

func installMmsrip() {
	// Método 1: Descargar y extraer .tar.gz
	mmsrip_tar_url := "https://nbenoit.tuxfamily.org/projects/mmsrip/mmsrip-0.7.0.tar.gz"
	mmsrip_cmd_tar := exec.Command("curl", "-L", mmsrip_tar_url, "-o", "package.tar.gz")
	err := mmsrip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmsrip_zip_url := "https://nbenoit.tuxfamily.org/projects/mmsrip/mmsrip-0.7.0.zip"
	mmsrip_cmd_zip := exec.Command("curl", "-L", mmsrip_zip_url, "-o", "package.zip")
	err = mmsrip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmsrip_bin_url := "https://nbenoit.tuxfamily.org/projects/mmsrip/mmsrip-0.7.0.bin"
	mmsrip_cmd_bin := exec.Command("curl", "-L", mmsrip_bin_url, "-o", "binary.bin")
	err = mmsrip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmsrip_src_url := "https://nbenoit.tuxfamily.org/projects/mmsrip/mmsrip-0.7.0.src.tar.gz"
	mmsrip_cmd_src := exec.Command("curl", "-L", mmsrip_src_url, "-o", "source.tar.gz")
	err = mmsrip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmsrip_cmd_direct := exec.Command("./binary")
	err = mmsrip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
