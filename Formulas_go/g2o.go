package main

// G2o - General framework for graph optimization
// Homepage: https://openslam-org.github.io/g2o.html

import (
	"fmt"
	
	"os/exec"
)

func installG2o() {
	// Método 1: Descargar y extraer .tar.gz
	g2o_tar_url := "https://github.com/RainerKuemmerle/g2o/archive/refs/tags/20230806_git.tar.gz"
	g2o_cmd_tar := exec.Command("curl", "-L", g2o_tar_url, "-o", "package.tar.gz")
	err := g2o_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	g2o_zip_url := "https://github.com/RainerKuemmerle/g2o/archive/refs/tags/20230806_git.zip"
	g2o_cmd_zip := exec.Command("curl", "-L", g2o_zip_url, "-o", "package.zip")
	err = g2o_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	g2o_bin_url := "https://github.com/RainerKuemmerle/g2o/archive/refs/tags/20230806_git.bin"
	g2o_cmd_bin := exec.Command("curl", "-L", g2o_bin_url, "-o", "binary.bin")
	err = g2o_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	g2o_src_url := "https://github.com/RainerKuemmerle/g2o/archive/refs/tags/20230806_git.src.tar.gz"
	g2o_cmd_src := exec.Command("curl", "-L", g2o_src_url, "-o", "source.tar.gz")
	err = g2o_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	g2o_cmd_direct := exec.Command("./binary")
	err = g2o_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
}
