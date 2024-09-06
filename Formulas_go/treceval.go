package main

// TrecEval - Evaluation software used in the Text Retrieval Conference
// Homepage: https://trec.nist.gov/

import (
	"fmt"
	
	"os/exec"
)

func installTrecEval() {
	// Método 1: Descargar y extraer .tar.gz
	treceval_tar_url := "https://github.com/usnistgov/trec_eval/archive/refs/tags/v9.0.8.tar.gz"
	treceval_cmd_tar := exec.Command("curl", "-L", treceval_tar_url, "-o", "package.tar.gz")
	err := treceval_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	treceval_zip_url := "https://github.com/usnistgov/trec_eval/archive/refs/tags/v9.0.8.zip"
	treceval_cmd_zip := exec.Command("curl", "-L", treceval_zip_url, "-o", "package.zip")
	err = treceval_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	treceval_bin_url := "https://github.com/usnistgov/trec_eval/archive/refs/tags/v9.0.8.bin"
	treceval_cmd_bin := exec.Command("curl", "-L", treceval_bin_url, "-o", "binary.bin")
	err = treceval_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	treceval_src_url := "https://github.com/usnistgov/trec_eval/archive/refs/tags/v9.0.8.src.tar.gz"
	treceval_cmd_src := exec.Command("curl", "-L", treceval_src_url, "-o", "source.tar.gz")
	err = treceval_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	treceval_cmd_direct := exec.Command("./binary")
	err = treceval_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
