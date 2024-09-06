package main

// Cmigemo - Migemo is a tool that supports Japanese incremental search with Romaji
// Homepage: https://www.kaoriya.net/software/cmigemo

import (
	"fmt"
	
	"os/exec"
)

func installCmigemo() {
	// Método 1: Descargar y extraer .tar.gz
	cmigemo_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmigemo/cmigemo-default-src-20110227.zip"
	cmigemo_cmd_tar := exec.Command("curl", "-L", cmigemo_tar_url, "-o", "package.tar.gz")
	err := cmigemo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmigemo_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmigemo/cmigemo-default-src-20110227.zip"
	cmigemo_cmd_zip := exec.Command("curl", "-L", cmigemo_zip_url, "-o", "package.zip")
	err = cmigemo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmigemo_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmigemo/cmigemo-default-src-20110227.zip"
	cmigemo_cmd_bin := exec.Command("curl", "-L", cmigemo_bin_url, "-o", "binary.bin")
	err = cmigemo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmigemo_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmigemo/cmigemo-default-src-20110227.zip"
	cmigemo_cmd_src := exec.Command("curl", "-L", cmigemo_src_url, "-o", "source.tar.gz")
	err = cmigemo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmigemo_cmd_direct := exec.Command("./binary")
	err = cmigemo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: nkf")
	exec.Command("latte", "install", "nkf").Run()
}
