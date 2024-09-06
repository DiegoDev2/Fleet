package main

// Diffstat - Produce graph of changes introduced by a diff file
// Homepage: https://invisible-island.net/diffstat/

import (
	"fmt"
	
	"os/exec"
)

func installDiffstat() {
	// Método 1: Descargar y extraer .tar.gz
	diffstat_tar_url := "https://invisible-mirror.net/archives/diffstat/diffstat-1.66.tgz"
	diffstat_cmd_tar := exec.Command("curl", "-L", diffstat_tar_url, "-o", "package.tar.gz")
	err := diffstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diffstat_zip_url := "https://invisible-mirror.net/archives/diffstat/diffstat-1.66.tgz"
	diffstat_cmd_zip := exec.Command("curl", "-L", diffstat_zip_url, "-o", "package.zip")
	err = diffstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diffstat_bin_url := "https://invisible-mirror.net/archives/diffstat/diffstat-1.66.tgz"
	diffstat_cmd_bin := exec.Command("curl", "-L", diffstat_bin_url, "-o", "binary.bin")
	err = diffstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diffstat_src_url := "https://invisible-mirror.net/archives/diffstat/diffstat-1.66.tgz"
	diffstat_cmd_src := exec.Command("curl", "-L", diffstat_src_url, "-o", "source.tar.gz")
	err = diffstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diffstat_cmd_direct := exec.Command("./binary")
	err = diffstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
