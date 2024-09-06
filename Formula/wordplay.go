package main

// Wordplay - Anagram generator
// Homepage: http://hsvmovies.com/static_subpages/personal_orig/wordplay/

import (
	"fmt"
	
	"os/exec"
)

func installWordplay() {
	// Método 1: Descargar y extraer .tar.gz
	wordplay_tar_url := "http://hsvmovies.com/static_subpages/personal_orig/wordplay/wordplay722.tar.Z"
	wordplay_cmd_tar := exec.Command("curl", "-L", wordplay_tar_url, "-o", "package.tar.gz")
	err := wordplay_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wordplay_zip_url := "http://hsvmovies.com/static_subpages/personal_orig/wordplay/wordplay722.tar.Z"
	wordplay_cmd_zip := exec.Command("curl", "-L", wordplay_zip_url, "-o", "package.zip")
	err = wordplay_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wordplay_bin_url := "http://hsvmovies.com/static_subpages/personal_orig/wordplay/wordplay722.tar.Z"
	wordplay_cmd_bin := exec.Command("curl", "-L", wordplay_bin_url, "-o", "binary.bin")
	err = wordplay_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wordplay_src_url := "http://hsvmovies.com/static_subpages/personal_orig/wordplay/wordplay722.tar.Z"
	wordplay_cmd_src := exec.Command("curl", "-L", wordplay_src_url, "-o", "source.tar.gz")
	err = wordplay_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wordplay_cmd_direct := exec.Command("./binary")
	err = wordplay_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
