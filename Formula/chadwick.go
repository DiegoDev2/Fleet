package main

// Chadwick - Tools for parsing Retrosheet MLB play-by-play files
// Homepage: https://chadwick.sourceforge.net/doc/index.html

import (
	"fmt"
	
	"os/exec"
)

func installChadwick() {
	// Método 1: Descargar y extraer .tar.gz
	chadwick_tar_url := "https://downloads.sourceforge.net/project/chadwick/chadwick-0.7/chadwick-0.7.2/chadwick-0.7.2.tar.gz"
	chadwick_cmd_tar := exec.Command("curl", "-L", chadwick_tar_url, "-o", "package.tar.gz")
	err := chadwick_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chadwick_zip_url := "https://downloads.sourceforge.net/project/chadwick/chadwick-0.7/chadwick-0.7.2/chadwick-0.7.2.zip"
	chadwick_cmd_zip := exec.Command("curl", "-L", chadwick_zip_url, "-o", "package.zip")
	err = chadwick_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chadwick_bin_url := "https://downloads.sourceforge.net/project/chadwick/chadwick-0.7/chadwick-0.7.2/chadwick-0.7.2.bin"
	chadwick_cmd_bin := exec.Command("curl", "-L", chadwick_bin_url, "-o", "binary.bin")
	err = chadwick_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chadwick_src_url := "https://downloads.sourceforge.net/project/chadwick/chadwick-0.7/chadwick-0.7.2/chadwick-0.7.2.src.tar.gz"
	chadwick_cmd_src := exec.Command("curl", "-L", chadwick_src_url, "-o", "source.tar.gz")
	err = chadwick_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chadwick_cmd_direct := exec.Command("./binary")
	err = chadwick_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
