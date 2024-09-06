package main

// Avimetaedit - Tool for embedding, validating, and exporting of AVI files metadata
// Homepage: https://mediaarea.net/AVIMetaEdit

import (
	"fmt"
	
	"os/exec"
)

func installAvimetaedit() {
	// Método 1: Descargar y extraer .tar.gz
	avimetaedit_tar_url := "https://mediaarea.net/download/binary/avimetaedit/1.0.2/AVIMetaEdit_CLI_1.0.2_GNU_FromSource.tar.bz2"
	avimetaedit_cmd_tar := exec.Command("curl", "-L", avimetaedit_tar_url, "-o", "package.tar.gz")
	err := avimetaedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avimetaedit_zip_url := "https://mediaarea.net/download/binary/avimetaedit/1.0.2/AVIMetaEdit_CLI_1.0.2_GNU_FromSource.tar.bz2"
	avimetaedit_cmd_zip := exec.Command("curl", "-L", avimetaedit_zip_url, "-o", "package.zip")
	err = avimetaedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avimetaedit_bin_url := "https://mediaarea.net/download/binary/avimetaedit/1.0.2/AVIMetaEdit_CLI_1.0.2_GNU_FromSource.tar.bz2"
	avimetaedit_cmd_bin := exec.Command("curl", "-L", avimetaedit_bin_url, "-o", "binary.bin")
	err = avimetaedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avimetaedit_src_url := "https://mediaarea.net/download/binary/avimetaedit/1.0.2/AVIMetaEdit_CLI_1.0.2_GNU_FromSource.tar.bz2"
	avimetaedit_cmd_src := exec.Command("curl", "-L", avimetaedit_src_url, "-o", "source.tar.gz")
	err = avimetaedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avimetaedit_cmd_direct := exec.Command("./binary")
	err = avimetaedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
