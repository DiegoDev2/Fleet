package main

// Nkf - Network Kanji code conversion Filter (NKF)
// Homepage: https://osdn.net/projects/nkf/

import (
	"fmt"
	
	"os/exec"
)

func installNkf() {
	// Método 1: Descargar y extraer .tar.gz
	nkf_tar_url := "https://dotsrc.dl.osdn.net/osdn/nkf/70406/nkf-2.1.5.tar.gz"
	nkf_cmd_tar := exec.Command("curl", "-L", nkf_tar_url, "-o", "package.tar.gz")
	err := nkf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nkf_zip_url := "https://dotsrc.dl.osdn.net/osdn/nkf/70406/nkf-2.1.5.zip"
	nkf_cmd_zip := exec.Command("curl", "-L", nkf_zip_url, "-o", "package.zip")
	err = nkf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nkf_bin_url := "https://dotsrc.dl.osdn.net/osdn/nkf/70406/nkf-2.1.5.bin"
	nkf_cmd_bin := exec.Command("curl", "-L", nkf_bin_url, "-o", "binary.bin")
	err = nkf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nkf_src_url := "https://dotsrc.dl.osdn.net/osdn/nkf/70406/nkf-2.1.5.src.tar.gz"
	nkf_cmd_src := exec.Command("curl", "-L", nkf_src_url, "-o", "source.tar.gz")
	err = nkf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nkf_cmd_direct := exec.Command("./binary")
	err = nkf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
