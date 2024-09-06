package main

// Wakeonlan - Sends magic packets to wake up network-devices
// Homepage: https://github.com/jpoliv/wakeonlan

import (
	"fmt"
	
	"os/exec"
)

func installWakeonlan() {
	// Método 1: Descargar y extraer .tar.gz
	wakeonlan_tar_url := "https://github.com/jpoliv/wakeonlan/archive/refs/tags/v0.42.tar.gz"
	wakeonlan_cmd_tar := exec.Command("curl", "-L", wakeonlan_tar_url, "-o", "package.tar.gz")
	err := wakeonlan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wakeonlan_zip_url := "https://github.com/jpoliv/wakeonlan/archive/refs/tags/v0.42.zip"
	wakeonlan_cmd_zip := exec.Command("curl", "-L", wakeonlan_zip_url, "-o", "package.zip")
	err = wakeonlan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wakeonlan_bin_url := "https://github.com/jpoliv/wakeonlan/archive/refs/tags/v0.42.bin"
	wakeonlan_cmd_bin := exec.Command("curl", "-L", wakeonlan_bin_url, "-o", "binary.bin")
	err = wakeonlan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wakeonlan_src_url := "https://github.com/jpoliv/wakeonlan/archive/refs/tags/v0.42.src.tar.gz"
	wakeonlan_cmd_src := exec.Command("curl", "-L", wakeonlan_src_url, "-o", "source.tar.gz")
	err = wakeonlan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wakeonlan_cmd_direct := exec.Command("./binary")
	err = wakeonlan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
