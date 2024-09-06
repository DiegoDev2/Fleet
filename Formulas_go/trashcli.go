package main

// TrashCli - Command-line interface to the freedesktop.org trashcan
// Homepage: https://github.com/andreafrancia/trash-cli

import (
	"fmt"
	
	"os/exec"
)

func installTrashCli() {
	// Método 1: Descargar y extraer .tar.gz
	trashcli_tar_url := "https://files.pythonhosted.org/packages/f7/6c/d51b36377c35e4f9e69af4d8b61a920f26251483cdc0165f5513da7aefeb/trash_cli-0.24.5.26.tar.gz"
	trashcli_cmd_tar := exec.Command("curl", "-L", trashcli_tar_url, "-o", "package.tar.gz")
	err := trashcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trashcli_zip_url := "https://files.pythonhosted.org/packages/f7/6c/d51b36377c35e4f9e69af4d8b61a920f26251483cdc0165f5513da7aefeb/trash_cli-0.24.5.26.zip"
	trashcli_cmd_zip := exec.Command("curl", "-L", trashcli_zip_url, "-o", "package.zip")
	err = trashcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trashcli_bin_url := "https://files.pythonhosted.org/packages/f7/6c/d51b36377c35e4f9e69af4d8b61a920f26251483cdc0165f5513da7aefeb/trash_cli-0.24.5.26.bin"
	trashcli_cmd_bin := exec.Command("curl", "-L", trashcli_bin_url, "-o", "binary.bin")
	err = trashcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trashcli_src_url := "https://files.pythonhosted.org/packages/f7/6c/d51b36377c35e4f9e69af4d8b61a920f26251483cdc0165f5513da7aefeb/trash_cli-0.24.5.26.src.tar.gz"
	trashcli_cmd_src := exec.Command("curl", "-L", trashcli_src_url, "-o", "source.tar.gz")
	err = trashcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trashcli_cmd_direct := exec.Command("./binary")
	err = trashcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
