package main

// Fprobe - Libpcap-based NetFlow probe
// Homepage: https://sourceforge.net/projects/fprobe/

import (
	"fmt"
	
	"os/exec"
)

func installFprobe() {
	// Método 1: Descargar y extraer .tar.gz
	fprobe_tar_url := "https://downloads.sourceforge.net/project/fprobe/fprobe/1.1/fprobe-1.1.tar.bz2"
	fprobe_cmd_tar := exec.Command("curl", "-L", fprobe_tar_url, "-o", "package.tar.gz")
	err := fprobe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fprobe_zip_url := "https://downloads.sourceforge.net/project/fprobe/fprobe/1.1/fprobe-1.1.tar.bz2"
	fprobe_cmd_zip := exec.Command("curl", "-L", fprobe_zip_url, "-o", "package.zip")
	err = fprobe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fprobe_bin_url := "https://downloads.sourceforge.net/project/fprobe/fprobe/1.1/fprobe-1.1.tar.bz2"
	fprobe_cmd_bin := exec.Command("curl", "-L", fprobe_bin_url, "-o", "binary.bin")
	err = fprobe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fprobe_src_url := "https://downloads.sourceforge.net/project/fprobe/fprobe/1.1/fprobe-1.1.tar.bz2"
	fprobe_cmd_src := exec.Command("curl", "-L", fprobe_src_url, "-o", "source.tar.gz")
	err = fprobe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fprobe_cmd_direct := exec.Command("./binary")
	err = fprobe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
