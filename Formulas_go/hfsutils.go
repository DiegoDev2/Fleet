package main

// Hfsutils - Tools for reading and writing Macintosh volumes
// Homepage: https://www.mars.org/home/rob/proj/hfs/

import (
	"fmt"
	
	"os/exec"
)

func installHfsutils() {
	// Método 1: Descargar y extraer .tar.gz
	hfsutils_tar_url := "https://ftp.osuosl.org/pub/clfs/conglomeration/hfsutils/hfsutils-3.2.6.tar.gz"
	hfsutils_cmd_tar := exec.Command("curl", "-L", hfsutils_tar_url, "-o", "package.tar.gz")
	err := hfsutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hfsutils_zip_url := "https://ftp.osuosl.org/pub/clfs/conglomeration/hfsutils/hfsutils-3.2.6.zip"
	hfsutils_cmd_zip := exec.Command("curl", "-L", hfsutils_zip_url, "-o", "package.zip")
	err = hfsutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hfsutils_bin_url := "https://ftp.osuosl.org/pub/clfs/conglomeration/hfsutils/hfsutils-3.2.6.bin"
	hfsutils_cmd_bin := exec.Command("curl", "-L", hfsutils_bin_url, "-o", "binary.bin")
	err = hfsutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hfsutils_src_url := "https://ftp.osuosl.org/pub/clfs/conglomeration/hfsutils/hfsutils-3.2.6.src.tar.gz"
	hfsutils_cmd_src := exec.Command("curl", "-L", hfsutils_src_url, "-o", "source.tar.gz")
	err = hfsutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hfsutils_cmd_direct := exec.Command("./binary")
	err = hfsutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
