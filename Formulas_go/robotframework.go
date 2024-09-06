package main

// RobotFramework - Open source test framework for acceptance testing
// Homepage: https://robotframework.org/

import (
	"fmt"
	
	"os/exec"
)

func installRobotFramework() {
	// Método 1: Descargar y extraer .tar.gz
	robotframework_tar_url := "https://files.pythonhosted.org/packages/51/84/3f1913910c8b877f13b444487861096049341a16b44a4aaee947e5546e2d/robotframework-7.0.1.zip"
	robotframework_cmd_tar := exec.Command("curl", "-L", robotframework_tar_url, "-o", "package.tar.gz")
	err := robotframework_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	robotframework_zip_url := "https://files.pythonhosted.org/packages/51/84/3f1913910c8b877f13b444487861096049341a16b44a4aaee947e5546e2d/robotframework-7.0.1.zip"
	robotframework_cmd_zip := exec.Command("curl", "-L", robotframework_zip_url, "-o", "package.zip")
	err = robotframework_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	robotframework_bin_url := "https://files.pythonhosted.org/packages/51/84/3f1913910c8b877f13b444487861096049341a16b44a4aaee947e5546e2d/robotframework-7.0.1.zip"
	robotframework_cmd_bin := exec.Command("curl", "-L", robotframework_bin_url, "-o", "binary.bin")
	err = robotframework_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	robotframework_src_url := "https://files.pythonhosted.org/packages/51/84/3f1913910c8b877f13b444487861096049341a16b44a4aaee947e5546e2d/robotframework-7.0.1.zip"
	robotframework_cmd_src := exec.Command("curl", "-L", robotframework_src_url, "-o", "source.tar.gz")
	err = robotframework_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	robotframework_cmd_direct := exec.Command("./binary")
	err = robotframework_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
