package main

// Httpstat - Curl statistics made simple
// Homepage: https://github.com/reorx/httpstat

import (
	"fmt"
	
	"os/exec"
)

func installHttpstat() {
	// Método 1: Descargar y extraer .tar.gz
	httpstat_tar_url := "https://github.com/reorx/httpstat/archive/refs/tags/1.3.2.tar.gz"
	httpstat_cmd_tar := exec.Command("curl", "-L", httpstat_tar_url, "-o", "package.tar.gz")
	err := httpstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpstat_zip_url := "https://github.com/reorx/httpstat/archive/refs/tags/1.3.2.zip"
	httpstat_cmd_zip := exec.Command("curl", "-L", httpstat_zip_url, "-o", "package.zip")
	err = httpstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpstat_bin_url := "https://github.com/reorx/httpstat/archive/refs/tags/1.3.2.bin"
	httpstat_cmd_bin := exec.Command("curl", "-L", httpstat_bin_url, "-o", "binary.bin")
	err = httpstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpstat_src_url := "https://github.com/reorx/httpstat/archive/refs/tags/1.3.2.src.tar.gz"
	httpstat_cmd_src := exec.Command("curl", "-L", httpstat_src_url, "-o", "source.tar.gz")
	err = httpstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpstat_cmd_direct := exec.Command("./binary")
	err = httpstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
