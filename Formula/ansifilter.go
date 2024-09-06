package main

// Ansifilter - Strip or convert ANSI codes into HTML, (La)Tex, RTF, or BBCode
// Homepage: http://andre-simon.de/doku/ansifilter/en/ansifilter.php

import (
	"fmt"
	
	"os/exec"
)

func installAnsifilter() {
	// Método 1: Descargar y extraer .tar.gz
	ansifilter_tar_url := "http://andre-simon.de/zip/ansifilter-2.20.tar.bz2"
	ansifilter_cmd_tar := exec.Command("curl", "-L", ansifilter_tar_url, "-o", "package.tar.gz")
	err := ansifilter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansifilter_zip_url := "http://andre-simon.de/zip/ansifilter-2.20.tar.bz2"
	ansifilter_cmd_zip := exec.Command("curl", "-L", ansifilter_zip_url, "-o", "package.zip")
	err = ansifilter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansifilter_bin_url := "http://andre-simon.de/zip/ansifilter-2.20.tar.bz2"
	ansifilter_cmd_bin := exec.Command("curl", "-L", ansifilter_bin_url, "-o", "binary.bin")
	err = ansifilter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansifilter_src_url := "http://andre-simon.de/zip/ansifilter-2.20.tar.bz2"
	ansifilter_cmd_src := exec.Command("curl", "-L", ansifilter_src_url, "-o", "source.tar.gz")
	err = ansifilter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansifilter_cmd_direct := exec.Command("./binary")
	err = ansifilter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
