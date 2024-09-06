package main

// Ndiff - Virtual package provided by nmap
// Homepage: https://www.math.utah.edu/~beebe/software/ndiff/

import (
	"fmt"
	
	"os/exec"
)

func installNdiff() {
	// Método 1: Descargar y extraer .tar.gz
	ndiff_tar_url := "https://ftp.math.utah.edu/pub/misc/ndiff-2.00.tar.gz"
	ndiff_cmd_tar := exec.Command("curl", "-L", ndiff_tar_url, "-o", "package.tar.gz")
	err := ndiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ndiff_zip_url := "https://ftp.math.utah.edu/pub/misc/ndiff-2.00.zip"
	ndiff_cmd_zip := exec.Command("curl", "-L", ndiff_zip_url, "-o", "package.zip")
	err = ndiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ndiff_bin_url := "https://ftp.math.utah.edu/pub/misc/ndiff-2.00.bin"
	ndiff_cmd_bin := exec.Command("curl", "-L", ndiff_bin_url, "-o", "binary.bin")
	err = ndiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ndiff_src_url := "https://ftp.math.utah.edu/pub/misc/ndiff-2.00.src.tar.gz"
	ndiff_cmd_src := exec.Command("curl", "-L", ndiff_src_url, "-o", "source.tar.gz")
	err = ndiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ndiff_cmd_direct := exec.Command("./binary")
	err = ndiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
