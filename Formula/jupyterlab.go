package main

// Jupyterlab - Interactive environments for writing and running code
// Homepage: https://jupyter.org/

import (
	"fmt"
	
	"os/exec"
)

func installJupyterlab() {
	// Método 1: Descargar y extraer .tar.gz
	jupyterlab_tar_url := "https://files.pythonhosted.org/packages/4a/78/ba006df6edaa561fe40be26c35e9da3f9316f071167cd7cc1a1a25bd2664/jupyterlab-4.2.5.tar.gz"
	jupyterlab_cmd_tar := exec.Command("curl", "-L", jupyterlab_tar_url, "-o", "package.tar.gz")
	err := jupyterlab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jupyterlab_zip_url := "https://files.pythonhosted.org/packages/4a/78/ba006df6edaa561fe40be26c35e9da3f9316f071167cd7cc1a1a25bd2664/jupyterlab-4.2.5.zip"
	jupyterlab_cmd_zip := exec.Command("curl", "-L", jupyterlab_zip_url, "-o", "package.zip")
	err = jupyterlab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jupyterlab_bin_url := "https://files.pythonhosted.org/packages/4a/78/ba006df6edaa561fe40be26c35e9da3f9316f071167cd7cc1a1a25bd2664/jupyterlab-4.2.5.bin"
	jupyterlab_cmd_bin := exec.Command("curl", "-L", jupyterlab_bin_url, "-o", "binary.bin")
	err = jupyterlab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jupyterlab_src_url := "https://files.pythonhosted.org/packages/4a/78/ba006df6edaa561fe40be26c35e9da3f9316f071167cd7cc1a1a25bd2664/jupyterlab-4.2.5.src.tar.gz"
	jupyterlab_cmd_src := exec.Command("curl", "-L", jupyterlab_src_url, "-o", "source.tar.gz")
	err = jupyterlab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jupyterlab_cmd_direct := exec.Command("./binary")
	err = jupyterlab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}
