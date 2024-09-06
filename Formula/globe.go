package main

// Globe - Prints ASCII graphic of currently-lit side of the Earth
// Homepage: https://web.archive.org/web/20230605040209/https://www.acme.com/software/globe/

import (
	"fmt"
	
	"os/exec"
)

func installGlobe() {
	// Método 1: Descargar y extraer .tar.gz
	globe_tar_url := "https://web.archive.org/web/20230605040209/https://www.acme.com/software/globe/globe_14Aug2014.tar.gz"
	globe_cmd_tar := exec.Command("curl", "-L", globe_tar_url, "-o", "package.tar.gz")
	err := globe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	globe_zip_url := "https://web.archive.org/web/20230605040209/https://www.acme.com/software/globe/globe_14Aug2014.zip"
	globe_cmd_zip := exec.Command("curl", "-L", globe_zip_url, "-o", "package.zip")
	err = globe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	globe_bin_url := "https://web.archive.org/web/20230605040209/https://www.acme.com/software/globe/globe_14Aug2014.bin"
	globe_cmd_bin := exec.Command("curl", "-L", globe_bin_url, "-o", "binary.bin")
	err = globe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	globe_src_url := "https://web.archive.org/web/20230605040209/https://www.acme.com/software/globe/globe_14Aug2014.src.tar.gz"
	globe_cmd_src := exec.Command("curl", "-L", globe_src_url, "-o", "source.tar.gz")
	err = globe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	globe_cmd_direct := exec.Command("./binary")
	err = globe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
