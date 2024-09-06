package main

// CvsFastExport - Export an RCS or CVS history as a fast-import stream
// Homepage: http://www.catb.org/~esr/cvs-fast-export/

import (
	"fmt"
	
	"os/exec"
)

func installCvsFastExport() {
	// Método 1: Descargar y extraer .tar.gz
	cvsfastexport_tar_url := "http://www.catb.org/~esr/cvs-fast-export/cvs-fast-export-1.68.tar.gz"
	cvsfastexport_cmd_tar := exec.Command("curl", "-L", cvsfastexport_tar_url, "-o", "package.tar.gz")
	err := cvsfastexport_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cvsfastexport_zip_url := "http://www.catb.org/~esr/cvs-fast-export/cvs-fast-export-1.68.zip"
	cvsfastexport_cmd_zip := exec.Command("curl", "-L", cvsfastexport_zip_url, "-o", "package.zip")
	err = cvsfastexport_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cvsfastexport_bin_url := "http://www.catb.org/~esr/cvs-fast-export/cvs-fast-export-1.68.bin"
	cvsfastexport_cmd_bin := exec.Command("curl", "-L", cvsfastexport_bin_url, "-o", "binary.bin")
	err = cvsfastexport_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cvsfastexport_src_url := "http://www.catb.org/~esr/cvs-fast-export/cvs-fast-export-1.68.src.tar.gz"
	cvsfastexport_cmd_src := exec.Command("curl", "-L", cvsfastexport_src_url, "-o", "source.tar.gz")
	err = cvsfastexport_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cvsfastexport_cmd_direct := exec.Command("./binary")
	err = cvsfastexport_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: asciidoctor")
	exec.Command("latte", "install", "asciidoctor").Run()
	fmt.Println("Instalando dependencia: cvs")
	exec.Command("latte", "install", "cvs").Run()
}
