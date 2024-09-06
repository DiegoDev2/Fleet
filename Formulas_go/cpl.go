package main

// Cpl - ISO-C libraries for developing astronomical data-reduction tasks
// Homepage: https://www.eso.org/sci/software/cpl/

import (
	"fmt"
	
	"os/exec"
)

func installCpl() {
	// Método 1: Descargar y extraer .tar.gz
	cpl_tar_url := "https://ftp.eso.org/pub/dfs/pipelines/libraries/cpl/cpl-7.3.2.tar.gz"
	cpl_cmd_tar := exec.Command("curl", "-L", cpl_tar_url, "-o", "package.tar.gz")
	err := cpl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpl_zip_url := "https://ftp.eso.org/pub/dfs/pipelines/libraries/cpl/cpl-7.3.2.zip"
	cpl_cmd_zip := exec.Command("curl", "-L", cpl_zip_url, "-o", "package.zip")
	err = cpl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpl_bin_url := "https://ftp.eso.org/pub/dfs/pipelines/libraries/cpl/cpl-7.3.2.bin"
	cpl_cmd_bin := exec.Command("curl", "-L", cpl_bin_url, "-o", "binary.bin")
	err = cpl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpl_src_url := "https://ftp.eso.org/pub/dfs/pipelines/libraries/cpl/cpl-7.3.2.src.tar.gz"
	cpl_cmd_src := exec.Command("curl", "-L", cpl_src_url, "-o", "source.tar.gz")
	err = cpl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpl_cmd_direct := exec.Command("./binary")
	err = cpl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cfitsio")
exec.Command("latte", "install", "cfitsio")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: wcslib")
exec.Command("latte", "install", "wcslib")
}
