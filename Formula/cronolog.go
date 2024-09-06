package main

// Cronolog - Web log rotation
// Homepage: https://web.archive.org/web/20140209202032/cronolog.org/

import (
	"fmt"
	
	"os/exec"
)

func installCronolog() {
	// Método 1: Descargar y extraer .tar.gz
	cronolog_tar_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/cronolog/cronolog-1.6.2.tar.gz"
	cronolog_cmd_tar := exec.Command("curl", "-L", cronolog_tar_url, "-o", "package.tar.gz")
	err := cronolog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cronolog_zip_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/cronolog/cronolog-1.6.2.zip"
	cronolog_cmd_zip := exec.Command("curl", "-L", cronolog_zip_url, "-o", "package.zip")
	err = cronolog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cronolog_bin_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/cronolog/cronolog-1.6.2.bin"
	cronolog_cmd_bin := exec.Command("curl", "-L", cronolog_bin_url, "-o", "binary.bin")
	err = cronolog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cronolog_src_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/cronolog/cronolog-1.6.2.src.tar.gz"
	cronolog_cmd_src := exec.Command("curl", "-L", cronolog_src_url, "-o", "source.tar.gz")
	err = cronolog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cronolog_cmd_direct := exec.Command("./binary")
	err = cronolog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
