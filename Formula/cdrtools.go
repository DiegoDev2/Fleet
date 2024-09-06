package main

// Cdrtools - CD/DVD/Blu-ray premastering and recording software
// Homepage: https://cdrtools.sourceforge.net/private/cdrecord.html

import (
	"fmt"
	
	"os/exec"
)

func installCdrtools() {
	// Método 1: Descargar y extraer .tar.gz
	cdrtools_tar_url := "https://downloads.sourceforge.net/project/cdrtools/alpha/cdrtools-3.02a09.tar.gz"
	cdrtools_cmd_tar := exec.Command("curl", "-L", cdrtools_tar_url, "-o", "package.tar.gz")
	err := cdrtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdrtools_zip_url := "https://downloads.sourceforge.net/project/cdrtools/alpha/cdrtools-3.02a09.zip"
	cdrtools_cmd_zip := exec.Command("curl", "-L", cdrtools_zip_url, "-o", "package.zip")
	err = cdrtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdrtools_bin_url := "https://downloads.sourceforge.net/project/cdrtools/alpha/cdrtools-3.02a09.bin"
	cdrtools_cmd_bin := exec.Command("curl", "-L", cdrtools_bin_url, "-o", "binary.bin")
	err = cdrtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdrtools_src_url := "https://downloads.sourceforge.net/project/cdrtools/alpha/cdrtools-3.02a09.src.tar.gz"
	cdrtools_cmd_src := exec.Command("curl", "-L", cdrtools_src_url, "-o", "source.tar.gz")
	err = cdrtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdrtools_cmd_direct := exec.Command("./binary")
	err = cdrtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: smake")
	exec.Command("latte", "install", "smake").Run()
}
