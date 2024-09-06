package main

// Dvdrtools - Fork of cdrtools DVD writer support
// Homepage: https://savannah.nongnu.org/projects/dvdrtools/

import (
	"fmt"
	
	"os/exec"
)

func installDvdrtools() {
	// Método 1: Descargar y extraer .tar.gz
	dvdrtools_tar_url := "https://savannah.nongnu.org/download/dvdrtools/dvdrtools-0.2.1.tar.gz"
	dvdrtools_cmd_tar := exec.Command("curl", "-L", dvdrtools_tar_url, "-o", "package.tar.gz")
	err := dvdrtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dvdrtools_zip_url := "https://savannah.nongnu.org/download/dvdrtools/dvdrtools-0.2.1.zip"
	dvdrtools_cmd_zip := exec.Command("curl", "-L", dvdrtools_zip_url, "-o", "package.zip")
	err = dvdrtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dvdrtools_bin_url := "https://savannah.nongnu.org/download/dvdrtools/dvdrtools-0.2.1.bin"
	dvdrtools_cmd_bin := exec.Command("curl", "-L", dvdrtools_bin_url, "-o", "binary.bin")
	err = dvdrtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dvdrtools_src_url := "https://savannah.nongnu.org/download/dvdrtools/dvdrtools-0.2.1.src.tar.gz"
	dvdrtools_cmd_src := exec.Command("curl", "-L", dvdrtools_src_url, "-o", "source.tar.gz")
	err = dvdrtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dvdrtools_cmd_direct := exec.Command("./binary")
	err = dvdrtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
