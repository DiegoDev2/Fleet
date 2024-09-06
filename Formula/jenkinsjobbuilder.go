package main

// JenkinsJobBuilder - Configure Jenkins jobs with YAML files stored in Git
// Homepage: https://docs.openstack.org/infra/jenkins-job-builder/

import (
	"fmt"
	
	"os/exec"
)

func installJenkinsJobBuilder() {
	// Método 1: Descargar y extraer .tar.gz
	jenkinsjobbuilder_tar_url := "https://files.pythonhosted.org/packages/c7/a1/0182d77739b546830015d8ff180e7287ed2f7f2533f0b4f98e1e371287e6/jenkins-job-builder-6.4.1.tar.gz"
	jenkinsjobbuilder_cmd_tar := exec.Command("curl", "-L", jenkinsjobbuilder_tar_url, "-o", "package.tar.gz")
	err := jenkinsjobbuilder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jenkinsjobbuilder_zip_url := "https://files.pythonhosted.org/packages/c7/a1/0182d77739b546830015d8ff180e7287ed2f7f2533f0b4f98e1e371287e6/jenkins-job-builder-6.4.1.zip"
	jenkinsjobbuilder_cmd_zip := exec.Command("curl", "-L", jenkinsjobbuilder_zip_url, "-o", "package.zip")
	err = jenkinsjobbuilder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jenkinsjobbuilder_bin_url := "https://files.pythonhosted.org/packages/c7/a1/0182d77739b546830015d8ff180e7287ed2f7f2533f0b4f98e1e371287e6/jenkins-job-builder-6.4.1.bin"
	jenkinsjobbuilder_cmd_bin := exec.Command("curl", "-L", jenkinsjobbuilder_bin_url, "-o", "binary.bin")
	err = jenkinsjobbuilder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jenkinsjobbuilder_src_url := "https://files.pythonhosted.org/packages/c7/a1/0182d77739b546830015d8ff180e7287ed2f7f2533f0b4f98e1e371287e6/jenkins-job-builder-6.4.1.src.tar.gz"
	jenkinsjobbuilder_cmd_src := exec.Command("curl", "-L", jenkinsjobbuilder_src_url, "-o", "source.tar.gz")
	err = jenkinsjobbuilder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jenkinsjobbuilder_cmd_direct := exec.Command("./binary")
	err = jenkinsjobbuilder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
