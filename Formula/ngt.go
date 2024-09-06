package main

// Ngt - Neighborhood graph and tree for indexing high-dimensional data
// Homepage: https://github.com/yahoojapan/NGT

import (
	"fmt"
	
	"os/exec"
)

func installNgt() {
	// Método 1: Descargar y extraer .tar.gz
	ngt_tar_url := "https://github.com/yahoojapan/NGT/archive/refs/tags/v2.2.4.tar.gz"
	ngt_cmd_tar := exec.Command("curl", "-L", ngt_tar_url, "-o", "package.tar.gz")
	err := ngt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ngt_zip_url := "https://github.com/yahoojapan/NGT/archive/refs/tags/v2.2.4.zip"
	ngt_cmd_zip := exec.Command("curl", "-L", ngt_zip_url, "-o", "package.zip")
	err = ngt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ngt_bin_url := "https://github.com/yahoojapan/NGT/archive/refs/tags/v2.2.4.bin"
	ngt_cmd_bin := exec.Command("curl", "-L", ngt_bin_url, "-o", "binary.bin")
	err = ngt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ngt_src_url := "https://github.com/yahoojapan/NGT/archive/refs/tags/v2.2.4.src.tar.gz"
	ngt_cmd_src := exec.Command("curl", "-L", ngt_src_url, "-o", "source.tar.gz")
	err = ngt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ngt_cmd_direct := exec.Command("./binary")
	err = ngt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
