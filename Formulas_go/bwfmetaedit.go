package main

// Bwfmetaedit - Tool for embedding, validating, and exporting BWF file metadata
// Homepage: https://mediaarea.net/BWFMetaEdit

import (
	"fmt"
	
	"os/exec"
)

func installBwfmetaedit() {
	// Método 1: Descargar y extraer .tar.gz
	bwfmetaedit_tar_url := "https://mediaarea.net/download/binary/bwfmetaedit/24.05/BWFMetaEdit_CLI_24.05_GNU_FromSource.tar.bz2"
	bwfmetaedit_cmd_tar := exec.Command("curl", "-L", bwfmetaedit_tar_url, "-o", "package.tar.gz")
	err := bwfmetaedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bwfmetaedit_zip_url := "https://mediaarea.net/download/binary/bwfmetaedit/24.05/BWFMetaEdit_CLI_24.05_GNU_FromSource.tar.bz2"
	bwfmetaedit_cmd_zip := exec.Command("curl", "-L", bwfmetaedit_zip_url, "-o", "package.zip")
	err = bwfmetaedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bwfmetaedit_bin_url := "https://mediaarea.net/download/binary/bwfmetaedit/24.05/BWFMetaEdit_CLI_24.05_GNU_FromSource.tar.bz2"
	bwfmetaedit_cmd_bin := exec.Command("curl", "-L", bwfmetaedit_bin_url, "-o", "binary.bin")
	err = bwfmetaedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bwfmetaedit_src_url := "https://mediaarea.net/download/binary/bwfmetaedit/24.05/BWFMetaEdit_CLI_24.05_GNU_FromSource.tar.bz2"
	bwfmetaedit_cmd_src := exec.Command("curl", "-L", bwfmetaedit_src_url, "-o", "source.tar.gz")
	err = bwfmetaedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bwfmetaedit_cmd_direct := exec.Command("./binary")
	err = bwfmetaedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
