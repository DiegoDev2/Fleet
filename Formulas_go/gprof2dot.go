package main

// Gprof2dot - Convert the output from many profilers into a Graphviz dot graph
// Homepage: https://github.com/jrfonseca/gprof2dot

import (
	"fmt"
	
	"os/exec"
)

func installGprof2dot() {
	// Método 1: Descargar y extraer .tar.gz
	gprof2dot_tar_url := "https://files.pythonhosted.org/packages/32/11/16fc5b985741378812223f2c6213b0a95cda333b797def622ac702d28e81/gprof2dot-2024.6.6.tar.gz"
	gprof2dot_cmd_tar := exec.Command("curl", "-L", gprof2dot_tar_url, "-o", "package.tar.gz")
	err := gprof2dot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gprof2dot_zip_url := "https://files.pythonhosted.org/packages/32/11/16fc5b985741378812223f2c6213b0a95cda333b797def622ac702d28e81/gprof2dot-2024.6.6.zip"
	gprof2dot_cmd_zip := exec.Command("curl", "-L", gprof2dot_zip_url, "-o", "package.zip")
	err = gprof2dot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gprof2dot_bin_url := "https://files.pythonhosted.org/packages/32/11/16fc5b985741378812223f2c6213b0a95cda333b797def622ac702d28e81/gprof2dot-2024.6.6.bin"
	gprof2dot_cmd_bin := exec.Command("curl", "-L", gprof2dot_bin_url, "-o", "binary.bin")
	err = gprof2dot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gprof2dot_src_url := "https://files.pythonhosted.org/packages/32/11/16fc5b985741378812223f2c6213b0a95cda333b797def622ac702d28e81/gprof2dot-2024.6.6.src.tar.gz"
	gprof2dot_cmd_src := exec.Command("curl", "-L", gprof2dot_src_url, "-o", "source.tar.gz")
	err = gprof2dot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gprof2dot_cmd_direct := exec.Command("./binary")
	err = gprof2dot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
}
