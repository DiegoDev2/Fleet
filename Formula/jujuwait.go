package main

// JujuWait - Juju plugin for waiting for deployments to settle
// Homepage: https://launchpad.net/juju-wait

import (
	"fmt"
	
	"os/exec"
)

func installJujuWait() {
	// Método 1: Descargar y extraer .tar.gz
	jujuwait_tar_url := "https://files.pythonhosted.org/packages/0c/2b/f4bd0138f941e4ba321298663de3f1c8d9368b75671b17aa1b8d41a154dc/juju-wait-2.8.4.tar.gz"
	jujuwait_cmd_tar := exec.Command("curl", "-L", jujuwait_tar_url, "-o", "package.tar.gz")
	err := jujuwait_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jujuwait_zip_url := "https://files.pythonhosted.org/packages/0c/2b/f4bd0138f941e4ba321298663de3f1c8d9368b75671b17aa1b8d41a154dc/juju-wait-2.8.4.zip"
	jujuwait_cmd_zip := exec.Command("curl", "-L", jujuwait_zip_url, "-o", "package.zip")
	err = jujuwait_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jujuwait_bin_url := "https://files.pythonhosted.org/packages/0c/2b/f4bd0138f941e4ba321298663de3f1c8d9368b75671b17aa1b8d41a154dc/juju-wait-2.8.4.bin"
	jujuwait_cmd_bin := exec.Command("curl", "-L", jujuwait_bin_url, "-o", "binary.bin")
	err = jujuwait_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jujuwait_src_url := "https://files.pythonhosted.org/packages/0c/2b/f4bd0138f941e4ba321298663de3f1c8d9368b75671b17aa1b8d41a154dc/juju-wait-2.8.4.src.tar.gz"
	jujuwait_cmd_src := exec.Command("curl", "-L", jujuwait_src_url, "-o", "source.tar.gz")
	err = jujuwait_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jujuwait_cmd_direct := exec.Command("./binary")
	err = jujuwait_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: juju")
	exec.Command("latte", "install", "juju").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
