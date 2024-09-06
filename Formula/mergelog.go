package main

// Mergelog - Merges httpd logs from web servers behind round-robin DNS
// Homepage: https://mergelog.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMergelog() {
	// Método 1: Descargar y extraer .tar.gz
	mergelog_tar_url := "https://downloads.sourceforge.net/project/mergelog/mergelog/4.5/mergelog-4.5.tar.gz"
	mergelog_cmd_tar := exec.Command("curl", "-L", mergelog_tar_url, "-o", "package.tar.gz")
	err := mergelog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mergelog_zip_url := "https://downloads.sourceforge.net/project/mergelog/mergelog/4.5/mergelog-4.5.zip"
	mergelog_cmd_zip := exec.Command("curl", "-L", mergelog_zip_url, "-o", "package.zip")
	err = mergelog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mergelog_bin_url := "https://downloads.sourceforge.net/project/mergelog/mergelog/4.5/mergelog-4.5.bin"
	mergelog_cmd_bin := exec.Command("curl", "-L", mergelog_bin_url, "-o", "binary.bin")
	err = mergelog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mergelog_src_url := "https://downloads.sourceforge.net/project/mergelog/mergelog/4.5/mergelog-4.5.src.tar.gz"
	mergelog_cmd_src := exec.Command("curl", "-L", mergelog_src_url, "-o", "source.tar.gz")
	err = mergelog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mergelog_cmd_direct := exec.Command("./binary")
	err = mergelog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
