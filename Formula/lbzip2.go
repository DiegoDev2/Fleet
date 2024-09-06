package main

// Lbzip2 - Parallel bzip2 utility
// Homepage: https://github.com/kjn/lbzip2

import (
	"fmt"
	
	"os/exec"
)

func installLbzip2() {
	// Método 1: Descargar y extraer .tar.gz
	lbzip2_tar_url := "https://web.archive.org/web/20170304050514/archive.lbzip2.org/lbzip2-2.5.tar.bz2"
	lbzip2_cmd_tar := exec.Command("curl", "-L", lbzip2_tar_url, "-o", "package.tar.gz")
	err := lbzip2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lbzip2_zip_url := "https://web.archive.org/web/20170304050514/archive.lbzip2.org/lbzip2-2.5.tar.bz2"
	lbzip2_cmd_zip := exec.Command("curl", "-L", lbzip2_zip_url, "-o", "package.zip")
	err = lbzip2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lbzip2_bin_url := "https://web.archive.org/web/20170304050514/archive.lbzip2.org/lbzip2-2.5.tar.bz2"
	lbzip2_cmd_bin := exec.Command("curl", "-L", lbzip2_bin_url, "-o", "binary.bin")
	err = lbzip2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lbzip2_src_url := "https://web.archive.org/web/20170304050514/archive.lbzip2.org/lbzip2-2.5.tar.bz2"
	lbzip2_cmd_src := exec.Command("curl", "-L", lbzip2_src_url, "-o", "source.tar.gz")
	err = lbzip2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lbzip2_cmd_direct := exec.Command("./binary")
	err = lbzip2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
