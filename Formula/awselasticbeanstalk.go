package main

// AwsElasticbeanstalk - Client for Amazon Elastic Beanstalk web service
// Homepage: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/eb-cli3.html

import (
	"fmt"
	
	"os/exec"
)

func installAwsElasticbeanstalk() {
	// Método 1: Descargar y extraer .tar.gz
	awselasticbeanstalk_tar_url := "https://files.pythonhosted.org/packages/ec/5e/96dbeec0f796ac7928f52ed61c6b3a44764ae4113185bb1e08fc4d758ba7/awsebcli-3.20.10.tar.gz"
	awselasticbeanstalk_cmd_tar := exec.Command("curl", "-L", awselasticbeanstalk_tar_url, "-o", "package.tar.gz")
	err := awselasticbeanstalk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awselasticbeanstalk_zip_url := "https://files.pythonhosted.org/packages/ec/5e/96dbeec0f796ac7928f52ed61c6b3a44764ae4113185bb1e08fc4d758ba7/awsebcli-3.20.10.zip"
	awselasticbeanstalk_cmd_zip := exec.Command("curl", "-L", awselasticbeanstalk_zip_url, "-o", "package.zip")
	err = awselasticbeanstalk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awselasticbeanstalk_bin_url := "https://files.pythonhosted.org/packages/ec/5e/96dbeec0f796ac7928f52ed61c6b3a44764ae4113185bb1e08fc4d758ba7/awsebcli-3.20.10.bin"
	awselasticbeanstalk_cmd_bin := exec.Command("curl", "-L", awselasticbeanstalk_bin_url, "-o", "binary.bin")
	err = awselasticbeanstalk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awselasticbeanstalk_src_url := "https://files.pythonhosted.org/packages/ec/5e/96dbeec0f796ac7928f52ed61c6b3a44764ae4113185bb1e08fc4d758ba7/awsebcli-3.20.10.src.tar.gz"
	awselasticbeanstalk_cmd_src := exec.Command("curl", "-L", awselasticbeanstalk_src_url, "-o", "source.tar.gz")
	err = awselasticbeanstalk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awselasticbeanstalk_cmd_direct := exec.Command("./binary")
	err = awselasticbeanstalk_cmd_direct.Run()
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
