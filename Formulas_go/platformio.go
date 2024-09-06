package main

// Platformio - Your Gateway to Embedded Software Development Excellence
// Homepage: https://platformio.org/

import (
	"fmt"
	
	"os/exec"
)

func installPlatformio() {
	// Método 1: Descargar y extraer .tar.gz
	platformio_tar_url := "https://files.pythonhosted.org/packages/13/78/8903f4f505a393ee48a18a00b4b9c866a726ef844d23ff3ce4863d710898/platformio-6.1.15.tar.gz"
	platformio_cmd_tar := exec.Command("curl", "-L", platformio_tar_url, "-o", "package.tar.gz")
	err := platformio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	platformio_zip_url := "https://files.pythonhosted.org/packages/13/78/8903f4f505a393ee48a18a00b4b9c866a726ef844d23ff3ce4863d710898/platformio-6.1.15.zip"
	platformio_cmd_zip := exec.Command("curl", "-L", platformio_zip_url, "-o", "package.zip")
	err = platformio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	platformio_bin_url := "https://files.pythonhosted.org/packages/13/78/8903f4f505a393ee48a18a00b4b9c866a726ef844d23ff3ce4863d710898/platformio-6.1.15.bin"
	platformio_cmd_bin := exec.Command("curl", "-L", platformio_bin_url, "-o", "binary.bin")
	err = platformio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	platformio_src_url := "https://files.pythonhosted.org/packages/13/78/8903f4f505a393ee48a18a00b4b9c866a726ef844d23ff3ce4863d710898/platformio-6.1.15.src.tar.gz"
	platformio_cmd_src := exec.Command("curl", "-L", platformio_src_url, "-o", "source.tar.gz")
	err = platformio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	platformio_cmd_direct := exec.Command("./binary")
	err = platformio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
