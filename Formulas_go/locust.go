package main

// Locust - Scalable user load testing tool written in Python
// Homepage: https://locust.io/

import (
	"fmt"
	
	"os/exec"
)

func installLocust() {
	// Método 1: Descargar y extraer .tar.gz
	locust_tar_url := "https://files.pythonhosted.org/packages/09/b3/af891e53130cff33329d05bd9e3632c125052d26981075659b86b8620599/locust-2.31.5.tar.gz"
	locust_cmd_tar := exec.Command("curl", "-L", locust_tar_url, "-o", "package.tar.gz")
	err := locust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	locust_zip_url := "https://files.pythonhosted.org/packages/09/b3/af891e53130cff33329d05bd9e3632c125052d26981075659b86b8620599/locust-2.31.5.zip"
	locust_cmd_zip := exec.Command("curl", "-L", locust_zip_url, "-o", "package.zip")
	err = locust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	locust_bin_url := "https://files.pythonhosted.org/packages/09/b3/af891e53130cff33329d05bd9e3632c125052d26981075659b86b8620599/locust-2.31.5.bin"
	locust_cmd_bin := exec.Command("curl", "-L", locust_bin_url, "-o", "binary.bin")
	err = locust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	locust_src_url := "https://files.pythonhosted.org/packages/09/b3/af891e53130cff33329d05bd9e3632c125052d26981075659b86b8620599/locust-2.31.5.src.tar.gz"
	locust_cmd_src := exec.Command("curl", "-L", locust_src_url, "-o", "source.tar.gz")
	err = locust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	locust_cmd_direct := exec.Command("./binary")
	err = locust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
}
