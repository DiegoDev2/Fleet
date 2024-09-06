package main

// Mkclean - Optimizes Matroska and WebM files
// Homepage: https://www.matroska.org/downloads/mkclean.html

import (
	"fmt"
	
	"os/exec"
)

func installMkclean() {
	// Método 1: Descargar y extraer .tar.gz
	mkclean_tar_url := "https://downloads.sourceforge.net/project/matroska/mkclean/mkclean-0.9.0.tar.bz2"
	mkclean_cmd_tar := exec.Command("curl", "-L", mkclean_tar_url, "-o", "package.tar.gz")
	err := mkclean_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkclean_zip_url := "https://downloads.sourceforge.net/project/matroska/mkclean/mkclean-0.9.0.tar.bz2"
	mkclean_cmd_zip := exec.Command("curl", "-L", mkclean_zip_url, "-o", "package.zip")
	err = mkclean_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkclean_bin_url := "https://downloads.sourceforge.net/project/matroska/mkclean/mkclean-0.9.0.tar.bz2"
	mkclean_cmd_bin := exec.Command("curl", "-L", mkclean_bin_url, "-o", "binary.bin")
	err = mkclean_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkclean_src_url := "https://downloads.sourceforge.net/project/matroska/mkclean/mkclean-0.9.0.tar.bz2"
	mkclean_cmd_src := exec.Command("curl", "-L", mkclean_src_url, "-o", "source.tar.gz")
	err = mkclean_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkclean_cmd_direct := exec.Command("./binary")
	err = mkclean_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
